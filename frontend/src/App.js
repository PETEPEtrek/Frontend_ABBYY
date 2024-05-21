import React from "react"
import Sidebar from "./components/Sidebar"
import MainContent from "./components/MainContent"
import TagsPanel from "./components/TagsPanel"
import SearchContextProvider from './search-context';
import FilterContextProvider from "./filter-context";

const App = () => {
    return (
        <SearchContextProvider>
            <FilterContextProvider>
                <div className="flex">
                    <Sidebar/>
                    <MainContent/>
                    <TagsPanel/>
                </div>
            </FilterContextProvider>
        </SearchContextProvider>
    );
}

export default App;
