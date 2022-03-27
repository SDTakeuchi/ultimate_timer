import { Html, Head, Main, NextScript } from 'next/document'
import Typography from '@material-ui/core/Typography';
import Box from '@material-ui/core/Box';

export default function Document() {
  return (
    <Html>
      <Head>
        <link
          rel="stylesheet"
          href="https://fonts.googleapis.com/css?family=Roboto:300,400,500,700&display=swap"
        />
      </Head>
      <body style={{backgroundColor: 'white', textAlign: 'center'}}>
        <Box className="container">
          <Typography variant="h2" component="h1">
            Ultimate Timer
          </Typography>
          <Main />
          <NextScript />
        </Box>
      </body>
    </Html>
  )
}