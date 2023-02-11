import React from 'react';
import axios from 'axios';

type Visit = {
    CreatedAt: string
    ID: number
    IP: string
    UserAgent: string
    VisitPath: {
        Path: string
    }
}

export default class VisitsIndexView extends React.Component {
    state = {
        visits: [] as Array<Visit>
    }

    componentDidMount() {
        this.indexVisits()
    }

    indexVisits() {
        axios.get(`https://api.yoku.dev/visit`)
            .then((res) => {
                this.setState({
                    visits: res.data.data,
                })
            })
            .catch((res) => {
                console.log(res)
            })
    }

    renderVisits(): Array<JSX.Element> {
        return this.state.visits.map(visit => {
            return (
                <tr
                    key={visit.ID}
                    className="border-b h-16"
                >
                    <td className="pr-8">{ new Date(visit.CreatedAt).toLocaleDateString([], {
                        year: 'numeric', month: 'long', day: 'numeric', hour: 'numeric', minute: 'numeric', second: 'numeric'
                    }) }</td>
                    <td className="pr-8">{ visit.IP }</td>
                    <td className="pr-8">
                        <a 
                            href={visit.VisitPath.Path}
                            target="_blank"
                            rel="noreferrer noopener"
                            className="underline hover:text-zinc-300"
                        >
                            { visit.VisitPath.Path }
                        </a>
                    </td>
                    <td className="pr-8">{ visit.UserAgent }</td>
                </tr>
            )
        })
        
    }

    render(): JSX.Element {
        return (
            <main className="p-4">
                <table className="text-left">
                    <thead>
                        <tr>
                            <th>Visited At</th>
                            <th>IP</th>
                            <th>Path</th>
                            <th>UserAgent</th>
                        </tr>
                    </thead>
                    <tbody>
                        { this.renderVisits() }
                    </tbody>
                </table>
            </main>
        )
    }
}
