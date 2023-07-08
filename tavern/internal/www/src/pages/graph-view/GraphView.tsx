import { DocumentNode, gql, useQuery } from "@apollo/client";
import React, { useState } from "react";
import { Link } from "react-router-dom";

import { CreateJobDrawer } from "../../components/create-job-drawer/CreateJobDrawer";
import { FormSteps } from "../../components/form-steps";
import { PageWrapper } from "../../components/page-wrapper";
import ForceGraph2D, { GraphData } from 'react-force-graph-2d';
import _ from "lodash";
const myData = {
    nodes: [
        {
            "id": "8589934593",
            "identifier": "AIj5GMrzjzNceQ6MjVJoxPswByTo2rNA0kI9XbZdBKt1TANXhH6B4Q/i1PrF3Mz5q2gzcgPNq0+BJrGrX9Q3Zg==",
            "name": "epic-practical-flameskull-8924169",
            "hostPlatform": "Unknown",
            "hostIdentifier": "zUBgPhPAM5R/JMcLozDisn/q6oujK7o56UDTxa0cnmFba/sjVgT2kyWIjra5CDh7Xrj6cIlPKo1/P+6CSFgukQ==",
            "hostPrimaryIP": ""
          },
          {
            "id": "8589934594",
            "identifier": "eBZEjCUhdhfi3XRDkRtW2Nqt+k9Mw7g0c0lTErLdyUJqLYpZc4lYw9zekCQD1VC/bwXo7IViStUumfwgkrBdpg==",
            "name": "upbeat-wizardly-night-hag-41778",
            "hostPlatform": "Unknown",
            "hostIdentifier": "zUBgPhPAM5R/JMcLozDisn/q6oujK7o56UDTxa0cnmFba/sjVgT2kyWIjra5CDh7Xrj6cIlPKo1/P+6CSFgukQ==",
            "hostPrimaryIP": ""
          },
          {
            "id": "8589934595",
            "identifier": "4MfH8TxO98rvYSdwCc20pwDrKXYUKPm2Yb9Pqzs0wZokWu4bYsTiz4laJpzraxKHxglx34d1BqONIgG9qhI+Qg==",
            "name": "spooky-celebrated-constrictor-snake-6635537",
            "hostPlatform": "Unknown",
            "hostIdentifier": "zUBgPhPAM5R/JMcLozDisn/q6oujK7o56UDTxa0cnmFba/sjVgT2kyWIjra5CDh7Xrj6cIlPKo1/P+6CSFgukQ==",
            "hostPrimaryIP": ""
          },    
    ],
    links: [
    ]
};

const GET_SESSIONS = gql`
    {
        sessions {
            __typename
            id
            identifier
            name
            hostPlatform
            hostIdentifier
            hostPrimaryIP
        }
    }
`;

const formatData = (data: any) => {
    const nodes = [] as any;
    const links = [] as any;
    data?.sessions.forEach((a: any) => {
        const b = Object.assign([], a);
        nodes.push(b)
    })
    return {
        nodes: _.uniqBy(nodes, 'id'),
        links
    };
}
  
export const GraphDiv = () => {
    const [graphData, setGraphData] = useState({ nodes: [] as Array<any>, links: [] as Array<any> });
    const { data } = useQuery(GET_SESSIONS, {
        onCompleted: (data) => setGraphData(formatData(data))
    });
    console.log(data);
    return <ForceGraph2D graphData={graphData} />;
    // return <div></div>
}

export const GraphView = () => {
    return (
        <PageWrapper>
            <div className="border-b border-gray-200 pb-5 sm:flex sm:items-center sm:justify-between">
                <h3 className="text-xl font-semibold leading-6 text-gray-900">Graph view</h3>
                <div className="mt-3 sm:mt-0 sm:ml-4">
                </div>
            </div>
            <GraphDiv />
        </PageWrapper>
    );
}