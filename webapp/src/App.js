import React from 'react';
import { BrowserRouter as Router, Route, Routes, Link } from 'react-router-dom';
import './App.css';
import Logo from './Logo';
import { Tech } from "./tech/Tech";
import { BlogList } from './blog/BlogList';
import { BlogDetail } from './blog/BlogDetail';
import { BlogForm } from './blog/BlogForm';

export function App() {
    return (
        <Router>
            <div className="app">
                <nav>
                    <h2 className="title">go-react-blog</h2>
                    <div className="logo"><Logo/></div>
                    <Link to="/">Home</Link> | 
                    <Link to="/blogs">Blogs</Link> |
                    <Link to="/blogs/new">New Blog</Link>
                </nav>

                <Routes>
                    <Route path="/" element={
                        <div>
                            <Tech/>
                        </div>
                    }/>
                    <Route path="/blogs" element={<BlogList/>}/>
                    <Route path="/blogs/new" element={<BlogForm/>}/>
                    <Route path="/blogs/:id" element={<BlogDetail/>}/>
                    <Route path="/blogs/:id/edit" element={<BlogForm/>}/>
                </Routes>
            </div>
        </Router>
    );
}
