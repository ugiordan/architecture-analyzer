import React, { useEffect, useState } from 'react';
import { Route, Switch } from 'react-router-dom';

interface User {
    id: number;
    name: string;
    bio: string;
}

const Dashboard: React.FC = () => {
    const [users, setUsers] = useState<User[]>([]);

    useEffect(() => {
        fetch('/api/users')
            .then(res => res.json())
            .then(data => setUsers(data));
    }, []);

    return (
        <div>
            {users.map(user => (
                <div key={user.id} dangerouslySetInnerHTML={{ __html: user.bio }} />
            ))}
        </div>
    );
};

const UserProfile: React.FC<{ userId: string }> = ({ userId }) => {
    const [user, setUser] = useState<User | null>(null);

    useEffect(() => {
        fetch(`/api/users/${userId}`).then(r => r.json()).then(setUser);
    }, [userId]);

    return user ? <h1>{user.name}</h1> : <p>Loading...</p>;
};

const AppRoutes = () => (
    <Switch>
        <Route path="/dashboard" component={Dashboard} />
        <Route path="/users/:id" component={UserProfile} />
    </Switch>
);

export default AppRoutes;
