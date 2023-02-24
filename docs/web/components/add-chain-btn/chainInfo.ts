import type { ChainInfo, BIP44 } from '@keplr-wallet/types';
import { Bech32Address } from '@keplr-wallet/cosmos';

const coinType60 = 60

const beraBip44: BIP44 = {
  coinType: coinType60,
};

const BERA = {
  coinDenom: 'bera',
  coinMinimalDenom: 'abera',
  coinDecimals: 18,
};

const BGT = {
  coinDenom: 'bgt',
  coinMinimalDenom: 'abgt',
  coinDecimals: 18,
};

const LOCAL_CHAIN_INFO: ChainInfo = {
  coinType: coinType60,
  rpc: 'http://localhost:26657',
  rest: 'http://localhost:1317',
  chainId: 'berachain_420-1',
  chainName: 'berachain-devnet',
  stakeCurrency: BGT,
  bip44: beraBip44,
  bech32Config: Bech32Address.defaultBech32Config('bera'),
  currencies: [BERA, BGT],
  feeCurrencies: [{
    ...BERA,
    gasPriceStep: {
      low: 10000000000,
      average: 25000000000,
      high: 40000000000,
    },
  }],
  features: ['ibc-transfer', 'ibc-go', 'eth-address-gen', 'eth-key-sign'],
};

export default LOCAL_CHAIN_INFO;