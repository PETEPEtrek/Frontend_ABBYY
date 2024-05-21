import React from "react"
import Sidebar from "./components/Sidebar"
import GamePage from "./components/GamePage"
const GamePageApp = () => {
    return (
        <div className="flex">
            <Sidebar/>
            <GamePage/>
        </div>
    );
}

export default GamePageApp;