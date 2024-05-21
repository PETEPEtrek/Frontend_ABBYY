import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';

const Login = () => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [redirect, setRedirect] = useState(false);
    const navigation = useNavigate();

    const submit = async (e) => {
        e.preventDefault();

        const response = await fetch('http://localhost:8000/user/login', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({
                email,
                password,
            }),
        });

        const content = await response.json();
        console.log(content)
        setRedirect(true);
    };

    if (redirect) {
        navigation('/');
    }

    return (
        <form onSubmit={submit}>
            <h1 className="h3 mb-3 fw-normal">Please sign in</h1>
            <input
                type="email"
                className="form-control"
                placeholder="Email address"
                required
                onChange={(e) => setEmail(e.target.value)}
            />

            <input
                type="password"
                className="form-control"
                placeholder="Password"
                required
                onChange={(e) => setPassword(e.target.value)}
            />

            <button className="w-100 btn btn-lg btn-primary" type="submit">
                Sign in
            </button>
        </form>
    );
};

export default Login;

