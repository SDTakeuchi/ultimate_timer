import * as React from 'react';
import Head from 'next/head';


export default (): JSX.Element => {
  const title: string = 'Ultimate Timer';
  const description: string = 
    'Ultimate Timer stores presets that you create. You don\'t need to get annoyed to set timer anymore.';
  const keyword: string = 'Ultimate Timer -any timer presets you want';

  return (
    <Head>
      <title>{title}</title>
      <meta property="og:title" content={title} />
      <meta property="og:description" content={description} />
      <meta name="keywords" content={keyword} />
      <meta property="og:type" content="blog" />
      <meta property="og:site_name" content={title} />
      <meta name="twitter:card" content="summary" />
      <meta name="twitter:site" content="@tcr_jp" />
      <meta name="twitter:title" content={title} />
      <meta name="twitter:description" content={description} />
      <link rel="shortcut icon" href={'https://t-cr.jp/favicon.ico'} />
      <link rel="apple-touch-icon" href={'https://t-cr.jp/logo.png'} />
    </Head>
  );
};
