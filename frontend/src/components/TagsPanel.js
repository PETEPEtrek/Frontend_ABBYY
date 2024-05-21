import React from "react"
import { useState, useContext } from "react"
import {GrCircleInformation} from "react-icons/gr"
import {GameTags} from "../data"
import {FilterContext} from "../filter-context";


const TagsPanel = () => {
    const [open2, setOpen2] = useState(false)

    const [adds, setAdds] = useState([false, false, false, false, false, false, false])

    const updateState = (index, newValue) => {
        const newStates = [...adds];
        newStates[index] = newValue;
        setAdds(newStates);
    };

    const filterContext = useContext(FilterContext);

    const [unselectedTags, setUnselectedTags] = useState(GameTags);
    const [selectedTags, setSelectedTags] = useState([]);

    const selectTagHandler = (tag, isChosen) => {
        const newStates = [...adds];
        newStates[tag.index] = isChosen;
        setAdds(newStates);
        let updatedSelectedTags = null;
        if (isChosen === true) {
            updatedSelectedTags = selectedTags.concat(tag);
            setSelectedTags(updatedSelectedTags);
            setUnselectedTags(
                unselectedTags.filter(unselectedTag => unselectedTag.title !== tag.title)
            );
            filterContext.tagSelector(updatedSelectedTags);
        } else {
            setUnselectedTags(unselectedTags.concat(tag));
            updatedSelectedTags = selectedTags.filter(
                selectedTag => selectedTag.title !== tag.title
            );
            setSelectedTags(updatedSelectedTags);
            filterContext.tagSelector(updatedSelectedTags);
        }
    };


    return (
        <div className={`bg-blue-600 h-screen p-5 pt-8 ${open2 ? "w-72" : "w-20"} duration-300 relative`}>
                <GrCircleInformation className="bg-white text-black text-3xl rounded-full absolute -left-3 top-9 border border-light-blue cursor-pointer" onClick={() => setOpen2(!open2)}/>
                <div className="inline-flex">
                    <h1 className={`text-white origin-left font-medium text-2xl duration-300 ${!open2 && "scale-0"}`}>Tags:</h1>
                </div>
                <ul className="pt-2">
                    {GameTags.map(tag => (
                        <>
                        <li key={tag.index} className={`text-white flex justify-between items-center text-sm gap-x-4 p-2 mt-2 duration-300 ${!open2 && "scale-0"}`}>
                            <span className="ml-2">{tag.title}</span>
                            <button className={`mr-2 text-2xl cursor-pointer hover:bg-gray-700 rounded-md ${adds[tag.index] ? "text-green-800" : "text-white"}`} onClick={() => selectTagHandler(tag, !adds[tag.index])}>+</button>
                        </li>
                        </>
                    ))}
                </ul>
        </div>
    );
}

export default TagsPanel;