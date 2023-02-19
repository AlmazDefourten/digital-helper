import { Link } from "react-router-dom";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import React from 'react';
import "./tabe.css";
import Project from "./Project";
import Application from "./Application";
import Grant from "./Grant";

function Tabe(){
    return(
        <Router>
            <div className="Tabb">
            <div className="Tabe">
                <div className="tab">
                    <Link to= "/project">
                        <div className="img_tab e"></div>
                        <p className="p_tab">МОИ ПРОЕКТЫ</p>
                    </Link>
                </div>
                <div className="tab">
                    <Link to="/application">
                        <div className="img_tab z"></div>
                        <p className="p_tab">МОИ ЗАЯВКИ</p>
                    </Link>
                </div>
                <div className="tab">
                    <Link to="/grant">
                        <div className="img_tab d"></div>
                        <p className="p_tab">МОИ ГРАНТЫ</p>
                    </Link>
                </div>
            </div>
            <div className="Pole">
                <Routes>
                    <Route path="/project" element={<Project />} />
                    <Route path="/application" element={<Application />} />
                    <Route path="/grant" element={<Grant />} />
                </Routes>
        </div>
        </div>
        </Router>
    );
}

export default Tabe;