const withNextra = require('nextra')({
  theme: 'nextra-theme-docs',
  themeConfig: './theme.config.jsx',
  defaultShowCopyCode: true,
  latex: true,

})

/** @type {import('next').NextConfig} */
const config = {
  exportPathMap: async function (
    defaultPathMap,
    { dev, dir, outDir, distDir, buildId }
  ) {
    return {
      '/': { page: '/docs' },
    }
  },
  reactStrictMode: true,
  images: {
    unoptimized: true
  },
  typescript: {
    // Disable type checking since eslint handles this
    ignoreBuildErrors: true,
  },
  async redirects() {
    return [
      {
        source: '/',
        destination: '/docs',
        permanent: false,
      }
    ];
  },
};
module.exports = withNextra(config)

// If you have other Next.js configurations, you can pass them as the parameter:
// module.exports = withNextra({ /* other next.js config */ })