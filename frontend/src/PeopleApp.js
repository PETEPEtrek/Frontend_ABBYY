import React from "react"
import PeopleSidebar from "./components/PeopleSidebar";
import PeopleMainContent from "./components/PeopleMainContent";
import PeopleTagsPanel from "./components/PeopleTagsPanel";
import SearchContextProvider from './search-context';
import FilterContextProvider from "./filter-context";

const PeopleApp = () => {
    return (
        <SearchContextProvider>
            <FilterContextProvider>
                <div className="flex">
                    <PeopleSidebar/>
                    <PeopleMainContent/>
                    <PeopleTagsPanel/>
                </div>
            </FilterContextProvider>
        </SearchContextProvider>
    );
}

export default PeopleApp;