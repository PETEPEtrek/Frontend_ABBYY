import React from "react"
import CharacterSidebar from "./components/CharacterSidebar"
import CharacterMainContent from "./components/CharacterMainContent"
import CharacterTagsPanel from "./components/CharacterTagsPanel"
import SearchContextProvider from './search-context';
import FilterContextProvider from "./filter-context";
const CharacterApp = () => {
    return (
        <SearchContextProvider>
            <FilterContextProvider>
                <div className="flex">
                    <CharacterSidebar/>
                    <CharacterMainContent/>
                    <CharacterTagsPanel/>
                </div>
            </FilterContextProvider>
        </SearchContextProvider>
    );
}

export default CharacterApp;