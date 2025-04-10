import React, { useState } from 'react'
import store from "../../assets/dummy.png"
import heart from "../../assets/heart.svg"

export default function Card() {

    const [showLikes, setShowLikes] = useState(false)

    return (
        <div 
            className="ImageDisplay__img"
            onMouseEnter={() => setShowLikes(true)}
            onMouseLeave={() => setShowLikes(false)}  
        >
            {showLikes && (
                <div
                  className="ImageDisplay__likes-display"
                >
                    <img src={heart}/>
                    <h3>255</h3>
                </div>
            )}
            <img src={store}/>
        </div>
    )
}
