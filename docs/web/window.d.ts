import { Window as KeplrWindow } from '@keplr-wallet/types';

declare global {
  // eslint-disable-next-line
  interface Window extends KeplrWindow {}
}
