import React, { useState, useEffect } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import axios from 'axios';
import './Blog.css';

export function BlogForm() {
    const [blog, setBlog] = useState({ title: '', body: '', coverURL: '' });
    const [error, setError] = useState(null);
    const { id } = useParams();
    const navigate = useNavigate();
    const isEdit = Boolean(id);

    useEffect(() => {
        if (isEdit) {
            const fetchBlog = async () => {
                try {
                    const response = await axios.get(`${process.env.REACT_APP_API_URL}/api/blogs/${id}`);
                    setBlog(response.data);
                } catch (err) {
                    setError('Failed to fetch blog');
                }
            };
            fetchBlog();
        }
    }, [id, isEdit]);

    const handleSubmit = async (e) => {
        e.preventDefault();
        try {
            if (isEdit) {
                await axios.put(`${process.env.REACT_APP_API_URL}/api/blogs/${id}`, blog);
            } else {
                await axios.post(`${process.env.REACT_APP_API_URL}/api/blogs`, blog);
            }
            navigate('/blogs');
        } catch (err) {
            setError('Failed to save blog');
        }
    };

    const handleChange = (e) => {
        setBlog({ ...blog, [e.target.name]: e.target.value });
    };

    return (
        <form className="blog-form" onSubmit={handleSubmit}>
            {error && <div className="error">{error}</div>}
            <div>
                <label>Title:</label>
                <input 
                    name="title" 
                    value={blog.title} 
                    onChange={handleChange} 
                    required 
                />
            </div>
            <div>
                <label>Cover URL:</label>
                <input 
                    name="coverURL" 
                    value={blog.coverURL} 
                    onChange={handleChange} 
                />
            </div>
            <div>
                <label>Content:</label>
                <textarea 
                    name="body" 
                    value={blog.body} 
                    onChange={handleChange} 
                    required 
                />
            </div>
            <button type="submit">{isEdit ? 'Update' : 'Create'} Blog</button>
        </form>
    );
}