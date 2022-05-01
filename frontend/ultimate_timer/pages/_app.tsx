import '../styles/globals.css'
import type { AppProps } from 'next/app'
import App from 'next/app'


function MyApp({ Component, pageProps }: AppProps) {
  return <Component {...pageProps} />
}
// App.getInitialProps = async () => ({ pageProps: {} })

export default MyApp
