// Copyright 2019 Google Cloud Platform Proxy Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// TODO(kyuc): Add unit tests

#include "src/envoy/http/backend_auth/filter.h"

#include <string>

#include "src/envoy/utils/filter_state_utils.h"

namespace Envoy {
namespace Extensions {
namespace HttpFilters {
namespace BackendAuth {

using ::google::protobuf::util::Status;
using Http::FilterDataStatus;
using Http::FilterHeadersStatus;
using Http::FilterTrailersStatus;
using Http::HeaderMap;
using Http::LowerCaseString;

namespace {
constexpr char kBearer[] = "Bearer ";
}  // namespace

FilterHeadersStatus Filter::decodeHeaders(HeaderMap& headers, bool) {
  absl::string_view operation = Utils::getStringFilterState(
      decoder_callbacks_->streamInfo().filterState(), Utils::kOperation);
  // NOTE: this shouldn't happen in practice because Path Matcher filter would
  // have already rejected the request.
  if (operation.empty()) {
    ENVOY_LOG(debug, "No operation found from DynamicMetadata");
    return FilterHeadersStatus::Continue;
  }

  ENVOY_LOG(debug, "Found operation: {}", operation);
  absl::string_view audience =
      config_->cfg_parser().getAudienceContext(operation);
  if (audience.empty()) {
    // JWT Token is not required for this operation.
    return FilterHeadersStatus::Continue;
  }

  const TokenSharedPtr jwt_token = config_->cfg_parser().getJwtToken(audience);
  if (!jwt_token) {
    ENVOY_LOG(debug, "Token not found for audience: {}", audience);
    return FilterHeadersStatus::Continue;
  }
  const auto& authorization = Http::Headers::get().Authorization;
  headers.remove(authorization);
  headers.addCopy(authorization, kBearer + *jwt_token);
  config_->stats().token_added_.inc();
  return FilterHeadersStatus::Continue;
}

FilterDataStatus Filter::decodeData(Buffer::Instance&, bool) {
  return FilterDataStatus::Continue;
}

FilterTrailersStatus Filter::decodeTrailers(HeaderMap&) {
  return FilterTrailersStatus::Continue;
}

void Filter::setDecoderFilterCallbacks(
    Http::StreamDecoderFilterCallbacks& callbacks) {
  decoder_callbacks_ = &callbacks;
}

}  // namespace BackendAuth
}  // namespace HttpFilters
}  // namespace Extensions
}  // namespace Envoy