import { DocsThemeConfig, useConfig } from 'nextra-theme-docs'
import { useRouter } from 'next/router'

export default {
    logo: <span>✨ Stargazer</span>,
    project: {
        link: 'https://github.com/berachain',
    },
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
        const { route } = useRouter()
        const socialCard =
            route === '/' || !title
                ? 'https://nextra.site/og.jpeg'
                : `https://nextra.site/api/og?title=${title}`

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
                <meta name="twitter:site:domain" content="nextra.site" />
                <meta name="twitter:url" content="https://nextra.site" />
                <meta
                    name="og:title"
                    content={title ? title + ' – Stargazer' : 'Stargazer'}
                />
                <meta name="og:image" content={socialCard} />
                <meta name="apple-mobile-web-app-title" content="Stargazer" />
                <link rel="icon" href="/favicon.svg" type="image/svg+xml" />
                <link rel="icon" href="/favicon.png" type="image/png" />
                <link
                    rel="icon"
                    href="/favicon-dark.svg"
                    type="image/svg+xml"
                    media="(prefers-color-scheme: dark)"
                />
                <link
                    rel="icon"
                    href="/favicon-dark.png"
                    type="image/png"
                    media="(prefers-color-scheme: dark)"
                />
            </>
        )
    },
    editLink: {
        text: 'Edit this page on GitHub →'
    },
    feedback: {
        content: 'Question? Give us feedback →',
        labels: 'feedback'
    },
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
                    © {new Date().getFullYear()} Berachain.
                </p>
            </div>
        )
    }
}