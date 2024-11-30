import React, { useState, useEffect } from 'react';
import { useParams, useNavigate, Link } from 'react-router-dom';
import axios from 'axios';
import './Blog.css';

export function BlogDetail() {
    const [blog, setBlog] = useState(null);
    const [error, setError] = useState(null);
    const { id } = useParams();
    const navigate = useNavigate();

    useEffect(() => {
        const fetchBlog = async () => {
            try {
                const response = await axios.get(`${process.env.REACT_APP_API_URL}/api/blogs/${id}`);
                setBlog(response.data);
            } catch (err) {
                setError('Failed to fetch blog');
            }
        };
        fetchBlog();
    }, [id]);

    const handleDelete = async () => {
        if (window.confirm('Are you sure you want to delete this blog?')) {
            try {
                await axios.delete(`${process.env.REACT_APP_API_URL}/api/blogs/${id}`);
                navigate('/blogs');
            } catch (err) {
                setError('Failed to delete blog');
            }
        }
    };

    if (error) return <div className="error">{error}</div>;
    if (!blog) return <div>Loading...</div>;

    return (
        <div className="blog-detail">
            <h2>{blog.title}</h2>
            {blog.coverURL && <img src={blog.coverURL} alt={blog.title} />}
            <p>{blog.body}</p>
            <div className="blog-actions">
                <Link to={`/blogs/${id}/edit`}>Edit</Link>
                <button onClick={handleDelete}>Delete</button>
            </div>
        </div>
    );
}