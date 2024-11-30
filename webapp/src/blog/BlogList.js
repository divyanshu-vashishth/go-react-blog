import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import axios from 'axios';
import './Blog.css';

export function BlogList() {
    const [blogs, setBlogs] = useState([]);
    const [error, setError] = useState(null);

    useEffect(() => {
        const fetchBlogs = async () => {
            try {
                const response = await axios.get(`${process.env.REACT_APP_API_URL}/api/blogs`);
                setBlogs(response.data);
            } catch (err) {
                setError('Failed to fetch blogs');
            }
        };
        fetchBlogs();
    }, []);

    if (error) return <div className="error">{error}</div>;

    return (
        <div className="blog-list">
            {blogs.map(blog => (
                <div key={blog.id} className="blog-card">
                    <h3>{blog.title}</h3>
                    {blog.coverURL && <img src={blog.coverURL} alt={blog.title} />}
                    <p>{blog.body.substring(0, 100)}...</p>
                    <Link to={`/blogs/${blog.id}`}>Read more</Link>
                </div>
            ))}
        </div>
    );
}