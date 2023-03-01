import type { ChainInfo, BIP44 } from '@keplr-wallet/types';
import { Bech32Address } from '@keplr-wallet/cosmos';

const coinType60 = 60

const gazerBip44: BIP44 = {
  coinType: coinType60,
};

const GAZER = {
  coinDenom: 'gazer',
  coinMinimalDenom: 'agazer',
  coinDecimals: 18,
};

const LOCAL_CHAIN_INFO: ChainInfo = {
  coinType: coinType60,
  rpc: 'http://localhost:26657',
  rest: 'http://localhost:1317',
  chainId: 'stargazer-2061',
  chainName: 'stargazer',
  stakeCurrency: GAZER,
  bip44: gazerBip44,
  bech32Config: Bech32Address.defaultBech32Config('stargazer'),
  currencies: [GAZER],
  feeCurrencies: [{
    ...GAZER,
    gasPriceStep: {
      low: 10000000000,
      average: 25000000000,
      high: 40000000000,
    },
  }],
  features: ['ibc-transfer', 'ibc-go', 'eth-address-gen', 'eth-key-sign'],
};

export default LOCAL_CHAIN_INFO;