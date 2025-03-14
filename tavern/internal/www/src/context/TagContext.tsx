import React, { createContext, useEffect } from "react";
import { ApolloError, gql, useQuery } from "@apollo/client";
import { TagContextType } from "../utils/consts";

const defaultValue = {data: undefined, isLoading: false, error: undefined} as {data: undefined | TagContextType, isLoading: boolean, error: ApolloError | undefined};

export const TagContext = createContext(defaultValue);
  
export const TagContextProvider = ({children}: {children: React.ReactNode}) => {

    const GET_TAG_FILTERS = gql`
        query GetSearchFilters($groupTag: TagWhereInput, $serviceTag: TagWhereInput){
            groupTags:tags(where: $groupTag) {
                id
                name
                kind   
            },
            serviceTags:tags(where: $serviceTag) {       
                id
                name
                kind   
            },
            beacons {
                id
                name
                principal
                lastSeenAt
                interval
                host{
                    name
                    primaryIP
                    platform
                    tags {
                        id
                        kind
                        name
                    }  
                }
            },
            hosts{
                id
                name
            }
        }
    `;
    const PARAMS = {
        variables: { 
            groupTag: { kind: "group" },
            serviceTag: { kind: "service" },
        }
    }
    const { loading: isLoading, error, data, startPolling, stopPolling } = useQuery(GET_TAG_FILTERS, PARAMS);

    useEffect(() => {
        startPolling(60000);
      return () => {
       stopPolling();
      }
    }, [startPolling, stopPolling])

  
    return (
      <TagContext.Provider value={{ data, isLoading, error }}>
        {children}
      </TagContext.Provider>
    );
};