import React from "react"
import Sidebar from "./components/Sidebar"
import PeoplePage from "./components/PeoplePage"
const PeoplePageApp = () => {
    return (
        <div className="flex">
            <Sidebar/>
            <PeoplePage/>
        </div>
    );
}

export default PeoplePageApp;