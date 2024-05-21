import React, {useState, useEffect, useContext} from "react"
import {FilterContext} from "../filter-context";
import {SearchContext} from "../search-context";

const PeopleMainContent =  () => {

    const [data, setData] = useState([]);
    const searchContext = useContext(SearchContext);
    const { tags } = useContext(FilterContext);

    useEffect(() => {
        const fetchData = async () => {
            const response = await fetch('http://localhost:8000/people');
            const jsonData = await response.json();
            setData(jsonData.data);
        };

        fetchData();
    }, []);


    let filteredArticles = null;
    if (data) {
        // 1. Tag Search
        filteredArticles = data.filter(article => {
            if (tags.length === 0) {
                return article;
            }
            if (tags.some(val => article.Tags.replace(/_/g, " ").includes(val.title))) {
                return article;
            } else {
                return null;
            }
        });
        filteredArticles = filteredArticles.filter(article => {
            if (article.Name.toLowerCase().includes(searchContext.query)) {
                return article;
            } else {
                return null;
            }
        });
    }


    return (
        <div className="flex-1 px-7 sm:px-5">
            <div className="flex justify-between items-center">
                <h1 className='text-2xl font-semibold'>VGDB - The Video Game Database</h1>
            </div>
            <div className="mb-10 sm:mb-0 mt-10 grid gap-4 grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4">
                {filteredArticles.map((people, index) => (
                    <a href={`/people/${people.ID}`} key={index} className="relative group bg-gray-900 py-10 sm:py-20 px-4 flex flex-col space-y-2 items-center cursor-pointer rounded-md hover:bg-gray-900/80 hover:smooth-hover">
                        <img className="w-20 h-20 object-cover object-center object-scale-down rounded-full" src={people.Image}/>
                        <h4 className="text-white text-2xl font-bold capitalize text-center">{people.Name}</h4>
                        <p className="text-white/50">{people.VoteNumber}</p>
                    </a>
                ))}
            </div>
        </div>
    );
}

export default PeopleMainContent;