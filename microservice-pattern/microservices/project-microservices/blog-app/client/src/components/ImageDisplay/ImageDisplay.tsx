import React, { useState } from 'react'
import "./ImageDisplay.css"
import { Container } from "react-bootstrap"
import Card from '../Card/Card'
import axios from "axios"
import {accessKeyId, secretAccessKey} from "../../keys"

function ImageDisplay() {

    const [file, setFile] = useState<File | null>(null);

    const savePhoto = async () => {

    }

    return (
        <div className="ImageDisplay">
            <input 
                type="file" 
                accept="image/*" 
                onChange={(e) => {
                    if(e.target.files){
                        setFile(e.target.files[0])
                    }
                }}
            />
            <button onClick={savePhoto}>Add Photo!</button>
            <Container className="ImageDisplay">
                <Card />
                <Card />
                <Card />
                <Card />
                {/* <Card />
                <Card />
                <Card />
                <Card />
                <Card /> */}
            </Container>
        </div>
    )
}

export default ImageDisplay
