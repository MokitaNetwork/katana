<!-- markdownlint-disable MD013 MD024 -->
<!-- markdown-link-check-disable -->

<!--
Changelog Guiding Principles:

Changelogs are for humans, not machines.
There should be an entry for every single version.
The same types of changes should be grouped.
Versions and sections should be linkable.
The latest version comes first.
The release date of each version is displayed.
Mention whether you follow Semantic Versioning.

Usage:

Change log entries are to be added to the Unreleased section under the
appropriate stanza (see below). Each entry should ideally include a tag and
the Github PR referenced in the following format:

* (<tag>) [#<PR-number>](https://github.com/mokitanetwork/katana/pull/<PR-number>) <changelog entry>

Types of changes (Stanzas):

Features: for new features.
Improvements: for changes in existing functionality.
Deprecated: for soon-to-be removed features.
Bug Fixes: for any bug fixes.
Client Breaking: for breaking Protobuf, CLI, gRPC and REST routes used by clients.
API Breaking: for breaking exported Go APIs used by developers.
State Machine Breaking: for any changes that result in a divergent application state.

To release a new version, ensure an appropriate release branch exists. Add a
release version and date to the existing Unreleased section which takes the form
of:

## [<version>](https://github.com/mokitanetwork/katana/releases/tag/<version>) - YYYY-MM-DD

Once the version is tagged and released, a PR should be made against the main
branch to incorporate the new changelog updates.

Ref: https://keepachangelog.com/en/1.0.0/
-->

# Changelog

## Unreleased

## [v2.0.1](https://github.com/mokitanetwork/katana/releases/tag/price-feeder/v2.0.1) 2022-12-01

### Bugs

- [1615](https://github.com/mokitanetwork/katana/pull/1615) Parse multiple candles from OsmosisV2 response
- [1635](https://github.com/mokitanetwork/katana/pull/1635) Vote on exchange rates even if one is missing.
- [1634](https://github.com/mokitanetwork/katana/pull/1634) Add minimum candle volume for low-trading assets.

### Improvements

- [1602](https://github.com/mokitanetwork/katana/pull/1602) Remove FTX provider.

## [v2.0.0](https://github.com/mokitanetwork/katana/releases/tag/price-feeder/v2.0.0) 2022-11-15

v2.0.0 of the price feeder contains numerous fixes for low-market-cap assets and API changes. It's highly recommended to switch to v2.0.0, especially as it removes the need to use the `ftx` provider for certain assets.

This was released as a part of [Katana Prop 27.](https://www.mintscan.io/katana/proposals/27)

### Bugs

- [1428](https://github.com/mokitanetwork/katana/pull/1428) Update katanad version to an actual tag.

### Features

- [1448](https://github.com/mokitanetwork/katana/pull/1448) Add crypto.com provider.
- [1496](https://github.com/mokitanetwork/katana/pull/1496) Dynamic provider minimum enforcement with CoinGecko API.
- [1510](https://github.com/mokitanetwork/katana/pull/1510) Integrate osmosis-api provider into price-feeder.
- [1534](https://github.com/mokitanetwork/katana/pull/1534) Query osmosis-api REST server for available asset pairs supported by it.
- [1554](https://github.com/mokitanetwork/katana/pull/1554) Convert remaining providers to the Websocket Controller.
- [1589](https://github.com/mokitanetwork/katana/pull/1589) Add Binance US provider.

### Improvements

- [1484](https://github.com/mokitanetwork/katana/pull/1484) Standardize websocket connection error for providers.
- [1509](https://github.com/mokitanetwork/katana/pull/1509) Update price feeder example config.
- [1527](https://github.com/mokitanetwork/katana/pull/1527) Update convertTickersToUSD and convertCandlesToUSD to public.

## [v1.0.0](https://github.com/mokitanetwork/katana/releases/tag/price-feeder%2Fv1.0.0) - 2022-09-19

### Features

- [1328](https://github.com/mokitanetwork/katana/pull/1328) Add bitget provider.
- [1339](https://github.com/mokitanetwork/katana/pull/1339) Add mexc provider.
- [1445](https://github.com/mokitanetwork/katana/pull/1445) Add computed prices api endpoints for debugging.

### Bugs

- [1338](https://github.com/mokitanetwork/katana/pull/1338) Fix websocket reconnections on remote closures.

## [v0.3.0](https://github.com/mokitanetwork/katana/releases/tag/price-feeder%2Fv0.3.0) - 2022-08-31

### Bugs

- [1084](https://github.com/mokitanetwork/katana/pull/1084) Initializes block height before subscription to fix an error message that appeared on the first few ticks.
- [1244](https://github.com/mokitanetwork/katana/pull/1244) Add verification for quote in conversion rate.
- [1264](https://github.com/mokitanetwork/katana/pull/1264) Convert osmosis candle timestamp from seconds to milliseconds.
- [1262](https://github.com/mokitanetwork/katana/pull/1262) Add verification for quote in tvwap map.
- [1268](https://github.com/mokitanetwork/katana/pull/1268) Don't panic when a provider has only out-of-date candles.
- [1291](https://github.com/mokitanetwork/katana/pull/1291) Set sdk version during build time.

### Improvements

- [#1121](https://github.com/mokitanetwork/katana/pull/1121) Use the cosmos-sdk telemetry package instead of our own.
- [#1032](https://github.com/mokitanetwork/katana/pull/1032) Update the accepted tvwap period from 3 minutes to 5 minutes.
- [#978](https://github.com/mokitanetwork/katana/pull/978) Cleanup the oracle package by moving deviation & conversion logic.
- [#1175](https://github.com/mokitanetwork/katana/pull/1175) Add type ProviderName.
- [#1255](https://github.com/mokitanetwork/katana/pull/1255) Move TickerPrice and CandlePrice to types package
- [#1374](https://github.com/mokitanetwork/katana/pull/1374) Add standard for telemetry metrics.
- [#1431](https://github.com/mokitanetwork/katana/pull/1431) Convert floats to sdk decimal using helper functions in all providers.
- [#1442](https://github.com/mokitanetwork/katana/pull/1442) Remove unnecessary method in recconection logic.

### Features

- [#1038](https://github.com/mokitanetwork/katana/pull/1038) Adds the option for validators to override API endpoints in our config.
- [#1002](https://github.com/mokitanetwork/katana/pull/1002) Add linting to the price feeder CI.
- [#1170](https://github.com/mokitanetwork/katana/pull/1170) Restrict price feeder quotes to USD, USDT, USDC, ETH, DAI, and BTC.
- [#1175](https://github.com/mokitanetwork/katana/pull/1175) Add ProviderName type to facilitate the reading of maps.
- [#1215](https://github.com/mokitanetwork/katana/pull/1215) Moved ProviderName to Name in provider package.
- [#1274](https://github.com/mokitanetwork/katana/pull/1274) Add option to set config by env variables.
- [#1299](https://github.com/mokitanetwork/katana/pull/1299) Add FTX as a provider.

## [v0.2.5](https://github.com/mokitanetwork/katana/releases/tag/price-feeder%2Fv0.2.5) - 2022-07-28

### Bugs

- [1177](https://github.com/mokitanetwork/katana/pull/1177) Update a deprecated osmosis api endpoint.

### Improvements

- [#1179](https://github.com/mokitanetwork/katana/pull/1179) Improve logs when unable to find prices.

## [v0.2.4](https://github.com/mokitanetwork/katana/releases/tag/price-feeder%2Fv0.2.4) - 2022-07-14

### Features

- [1110](https://github.com/mokitanetwork/katana/pull/1110) Add the ability to detect deviations with multi-quoted prices, ex. using BTC/USD and BTC/ETH at the same time.
- [#998](https://github.com/mokitanetwork/katana/pull/998) Make deviation thresholds configurable for stablecoin support.

## [v0.2.3](https://github.com/mokitanetwork/katana/releases/tag/price-feeder%2Fv0.2.3) - 2022-06-30

### Improvements

- [#1069](https://github.com/mokitanetwork/katana/pull/1069) Subscribe to node event EventNewBlockHeader to have the current chain height.

## [v0.2.2](https://github.com/mokitanetwork/katana/releases/tag/price-feeder%2Fv0.2.2) - 2022-06-27

### Improvements

- [#1050](https://github.com/mokitanetwork/katana/pull/1050) Cache x/oracle params to decrease the number of queries to nodes.

### Features

- [#925](https://github.com/mokitanetwork/katana/pull/925) Require stablecoins to be converted to USD to protect against depegging.

## [v0.2.1](https://github.com/mokitanetwork/katana/releases/tag/price-feeder%2Fv0.2.1) - 2022-04-06

### Improvements

- [#766](https://github.com/mokitanetwork/katana/pull/766) Update deps to use katana v2.0.0.

## [v0.2.0](https://github.com/mokitanetwork/katana/releases/tag/price-feeder%2Fv0.2.0) - 2022-04-04

### Features

- [#730](https://github.com/mokitanetwork/katana/pull/730) Update the mock provider to use a new spreadsheet which uses randomness.

### Improvements

- [#684](https://github.com/mokitanetwork/katana/pull/684) Log errors when providers are unable to unmarshal candles and tickers, instead of either one.
- [#732](https://github.com/mokitanetwork/katana/pull/732) Set oracle functions to public to facilitate usage in other repositories.

### Bugs

- [#732](https://github.com/mokitanetwork/katana/pull/732) Fixes an issue where filtering out erroneous providers' candles wasn't working.

## [v0.1.4](https://github.com/mokitanetwork/katana/releases/tag/price-feeder%2Fv0.1.4) - 2022-03-24

### Features

- [#648](https://github.com/mokitanetwork/katana/pull/648) Add Coinbase as a provider.
- [#679](https://github.com/mokitanetwork/katana/pull/679) Add a configurable provider timeout, which defaults to 100ms.

### Bug Fixes

- [#675](https://github.com/mokitanetwork/katana/pull/675) Add necessary input validation to SubscribePairs in the price feeder.

## [v0.1.3](https://github.com/mokitanetwork/katana/releases/tag/price-feeder%2Fv0.1.3) - 2022-03-21

### Features

- [#649](https://github.com/mokitanetwork/katana/pull/649) Add "GetAvailablePairs" function to providers.

## [v0.1.2](https://github.com/mokitanetwork/katana/releases/tag/price-feeder%2Fv0.1.2) - 2022-03-08

### Features

- [#592](https://github.com/mokitanetwork/katana/pull/592) Add subscribe ticker function to the following providers: Binance, Huobi, Kraken, and Okx.
- [#601](https://github.com/mokitanetwork/katana/pull/601) Use TVWAP formula for determining prices when available.
- [#609](https://github.com/mokitanetwork/katana/pull/609) TVWAP faulty provider detection.

### Bug Fixes

- [#607](https://github.com/mokitanetwork/katana/pull/607) Fix kraken provider timestamp unit.

### Refactor

- [#610](https://github.com/mokitanetwork/katana/pull/610) Split subscription of ticker and candle channels.

## [v0.1.1](https://github.com/mokitanetwork/katana/releases/tag/price-feeder%2Fv0.1.1) - 2022-03-01

### Features

- [#502](https://github.com/mokitanetwork/katana/pull/502) Faulty provider detection: discard prices that are not within 2𝜎 of others.
- [#536](https://github.com/mokitanetwork/katana/pull/536) Force a minimum of three providers per asset.
- [#522](https://github.com/mokitanetwork/katana/pull/522) Add Okx as a provider.
- [#551](https://github.com/mokitanetwork/katana/pull/551) Update Binance provider to use WebSocket.
- [#569](https://github.com/mokitanetwork/katana/pull/569) Update Huobi provider to use WebSocket.
- [#540](https://github.com/mokitanetwork/katana/pull/536) Use environment vars / standard input for the keyring password instead of the config file.
- [#580](https://github.com/mokitanetwork/katana/pull/580) Update Kraken provider to use WebSocket.

### Bug Fixes

- [#552](https://github.com/mokitanetwork/katana/pull/552) Stop requiring telemetry during config validation.
- [#573](https://github.com/mokitanetwork/katana/pull/573) Strengthen CORS settings.
- [#574](https://github.com/mokitanetwork/katana/pull/574) Stop registering metrics endpoint if telemetry is disabled.

### Refactor

- [#587](https://github.com/mokitanetwork/katana/pull/587) Clean up logs from price feeder providers.

## [v0.1.0](https://github.com/mokitanetwork/katana/releases/tag/price-feeder%2Fv0.1.0) - 2022-02-07

### Features

- Initial release!!!
