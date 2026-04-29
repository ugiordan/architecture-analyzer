import express from 'express';

function handleRequest(req: express.Request, res: express.Response): void {
    const body = req.body;
    const payload = JSON.parse(body);
    const name: string = payload.name;
    const query = "SELECT * FROM users WHERE name = " + name;
    res.send(query);
}

function simpleAssignment(): string {
    const x = "hello";
    const y = x;
    return y;
}

function fieldAccess(user: { email: string }): string {
    const email = user.email;
    return email;
}

const arrowFn = (data: string): string => {
    const result = data.trim();
    return result;
};
