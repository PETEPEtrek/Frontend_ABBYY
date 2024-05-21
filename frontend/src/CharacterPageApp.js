import React from "react"
import Sidebar from "./components/Sidebar"
import CharacterPage from "./components/CharacterPage"
const CharacterPageApp = () => {
    return (
        <div className="flex">
            <Sidebar/>
            <CharacterPage/>
        </div>
    );
}

export default CharacterPageApp;