import { DocsThemeConfig, useConfig } from 'nextra-theme-docs'
import { useRouter } from 'next/router'

export default {
    logo: <span>✨ Stargazer</span>,
    logoLink: '/docs',
    project: {
        link: 'https://github.com/berachain/stargazer',
    },
    docsRepositoryBase: "http://github.com/berachain/stargazer",
    useNextSeoProps() {
        const { route } = useRouter()
        if (route !== '/') {
            return {
                titleTemplate: '%s – Stargazer Docs'
            }
        }
    },
    head: function useHead() {
        const { title } = useConfig()
        const socialCard = '/header.jpeg'

        return (
            <>
                <meta name="msapplication-TileColor" content="#fff" />
                <meta name="theme-color" content="#fff" />
                <meta name="viewport" content="width=device-width, initial-scale=1.0" />
                <meta httpEquiv="Content-Language" content="en" />
                <meta
                    name="description"
                    content="Stargazer brings EVM to Cosmos in a new way"
                />
                <meta
                    name="og:description"
                    content="Stargazer brings EVM to Cosmos in a new way"
                />
                <meta name="twitter:card" content="summary_large_image" />
                <meta name="twitter:image" content={socialCard} />
                <meta name="twitter:site:domain" content="berachain.com" />
                <meta name="twitter:url" content="https://berachain.com/" />
                <meta
                    name="og:title"
                    content={title ? title + ' – Stargazer' : 'Stargazer'}
                />
                <meta name="og:image" content={socialCard} />
                <meta name="apple-mobile-web-app-title" content="Stargazer" />
                <link rel="icon" href="/berachain.svg" type="image/svg+xml" />
                <link rel="icon" href="/berachain.png" type="image/png" />
                <link rel="icon" href="/berachain.ico"/>
                <link
                    rel="icon"
                    href="/berachain.svg"
                    type="image/svg+xml"
                    media="(prefers-color-scheme: dark)"
                />
                <link
                    rel="icon"
                    href="/berachain.png"
                    type="image/png"
                    media="(prefers-color-scheme: dark)"
                />
            </>
        )
    },
    editLink: {
        text: 'Edit this page on GitHub →'
    },
    feedback: false,
    sidebar: {
        titleComponent({ title, type }) {
            if (type === 'separator') {
                return <span className="cursor-default">{title}</span>
            }
            return <>{title}</>
        },
        defaultMenuCollapseLevel: 1,
        toggleButton: true,
    },
    footer: {
        text: (
            <div>
                <p>
                    © {new Date().getFullYear()} Berachain Foundation.
                </p>
            </div>
        )
    }
}