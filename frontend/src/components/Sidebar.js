import React from "react"
import { useState, useContext} from "react"
import {SearchContext} from "../search-context"
import { BsArrowLeftShort , BsSearch} from "react-icons/bs"
import {IoGameControllerOutline} from "react-icons/io5"
import { useLocation } from "react-router-dom";
import { UserButton } from "@clerk/clerk-react"
const Sidebar = () => {
    const [open, setOpen] = useState(true);
    const location = useLocation();
    const { query, searchHandler } = useContext(SearchContext);

    const Menus = [
        {index: 0, title: "Games", ref: "/"},
        {index: 1, title: "People", ref: "/people"},
        {index: 2, title: "Characters", ref: "/characters"},
        {index: 3, title: "Log in", ref: "/user/login"},
        {index: 4, title: "Sign in", ref: "/user/sign_in"}
    ]

    return (
        <div className={`bg-blue-600 h-screen p-5 pt-8 ${open ? "w-72" : "w-20"} duration-300 relative`}>
                <BsArrowLeftShort className={`bg-white text-black text-3xl rounded-full absolute -right-3 top-9 border border-light-blue cursor-pointer ${!open && "rotate-180"}`} onClick={() => setOpen(!open)}/>
                <div className="inline-flex">
                    <IoGameControllerOutline className='bg-blue-600 text-4xl rounded cursor-pointer block float-left mr-2'/>
                    <h1 className={`text-white origin-left font-medium text-2xl duration-300 ${!open && "scale-0"}`}>VGDB</h1>
                </div>
                <div className={`flex items-center rounded-md bg-gray-400 mt-6 ${!open ? "px-2.5": "px-4"} py-2`}>
                    <BsSearch className={`text-black text-lg block float-left cursor-pointer ${open && "mr-2"}`}/>
                    <input type={"search"} placeholder="Search" onChange={e => searchHandler(e.target.value)}
                           value={query} className={`text-base bg-transparent w-full text-black focus:outline-none ${!open && "hidden"}`}/>
                </div>
                <ul className="pt-2">
                    {Menus.map(menu => (
                        <>
                        {menu.title === "Log in" && location.pathname === "/protected" ? (
                        <li key={menu.index} className={`text-white-300 text-sm flex-item-center gap-x-4 cursor-pointer p-2 hover:bg-gray-400 rounded-md mt-2 duration-300 ${!open && "scale-0"}`}>
                            <UserButton />
                        </li>
                        ) : (
                        <li key={menu.index} className={`text-white-300 text-sm flex-item-center gap-x-4 cursor-pointer p-2 hover:bg-gray-400 rounded-md mt-2 duration-300 ${!open && "scale-0"}`}>
                            <a href={menu.ref}>{menu.title}</a>
                        </li>
                        )}
                        </>
                    ))}
                </ul>
        </div>
    );
}


export default Sidebar;