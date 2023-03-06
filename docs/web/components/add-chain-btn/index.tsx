import LOCAL_CHAIN_INFO from './chainInfo'
import styles from './style.module.css'

const addNetwork = async () => {
    if(!window.keplr) {
        alert('Please install keplr')
        return
    }
    try {
        await window.keplr.experimentalSuggestChain(LOCAL_CHAIN_INFO);
    } catch(e) {
        return
    }
    return
}
export function AddKeplrBtn() {
  return (
    <a onClick={async() => {await addNetwork()}} className={styles.cta}>
        Add Local Network
    </a>
  )
}
