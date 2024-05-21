import React, {useEffect, useState} from "react";
import { useParams} from "react-router-dom";
import styled from "styled-components";
const PeoplePage =  () => {
    const {id} = useParams()

    const [showMore, setShowMore] = React.useState(false)
    const [data, setData] = useState([]);

    useEffect(() => {
        const fetchData = async () => {
            const response = await fetch(`http://localhost:8000/people/${id}`);
            const jsonData = await response.json();
            setData(jsonData.data);
        };

        fetchData();
    }, []);


    return (
        <AnimeItemStyled>
            <div>
                <h1>{data.Name}</h1>
                <div className="details">
                    <div className="detail">
                        <div className="image">
                            <img src={data.Image} alt=""></img>
                        </div>
                        <div className="anime-details">
                            <p flex gap-4><span font-semibold text-slate-950>Date:</span><span>{data.BirthDate}</span></p>
                            <p flex gap-4><span font-semibold text-slate-950>Tags:</span><span>{String(data.Tags).replace(/\//g, ",").replace(/_/g, " ").substring(1, String(data.Tags).length - 1)}</span></p>
                            <p flex gap-4><span font-semibold text-slate-950>Rating:</span><span>{data.Score}</span></p>
                        </div>
                    </div>
                    <p className="description">
                        {showMore ? data.Story : data.Story?.substring(0, 450) + "..."}
                        <button
                            onClick={() => {
                                setShowMore(!showMore)
                            }}>{showMore ? 'Show Less' : "Read More"}</button>
                    </p>
                </div>
            </div>
        </AnimeItemStyled>
    )
}

const AnimeItemStyled = styled.div`
    padding: 3rem 18rem;
    background-color: #EDEDED;
    h1{
        display: inline-block;
        font-size: 5rem;
        margin-bottom: 2rem;
    }
    .title{
        display: inline-block;
        margin: 3rem 0;
        font-size: 2rem;
    }

    .description{
        margin-top: 2rem;
        color: #6c7983;
        line-height: 1.7rem;
        button{
            background-color: transparent;
            border: none;
            outline: none;
            cursor: pointer;
            font-size: 1.2rem;
            color: #27AE60;
            font-weight: 600;
        }
    }

    .trailer-con{
        display: flex;
        justify-content: center;
        align-items: center;
        iframe{
            outline: none;
            border: 5px solid #e5e7eb;
            padding: 1.5rem;
            border-radius: 10px;
            background-color: #FFFFFF;
        }
    }

    .details{
        background-color: #fff;
        border-radius: 20px;
        padding: 2rem;
        border: 5px solid #e5e7eb;
        .detail{
            display: grid;
            grid-template-columns: repeat(2, 1fr);
            gap: 4rem;
            img{
                border-radius: 7px;
            }
        }
        .anime-details{
            display: flex;
            flex-direction: column;
            justify-content: space-between;
            p{
                display: flex;
                gap: 4rem;
            }
            p span:first-child{
                font-weight: 600;
                color: #454e56;
            }
        }
    }

`;
export default PeoplePage;