import express, { Request, Response, NextFunction } from 'express';
import { Pool } from 'pg';

const app = express();
const pool = new Pool();
const DB_PASSWORD = process.env.DB_PASSWORD;

function authMiddleware(req: Request, res: Response, next: NextFunction): void {
    const token = req.headers.authorization;
    if (!token) {
        res.status(401).send('Unauthorized');
        return;
    }
    next();
}

app.get('/users', authMiddleware, async (req: Request, res: Response) => {
    const result = await pool.query(`SELECT * FROM users WHERE active = ${req.query.active}`);
    res.json(result.rows);
});

app.post('/users', async (req: Request, res: Response) => {
    const { name, email } = req.body;
    await pool.query('INSERT INTO users (name, email) VALUES ($1, $2)', [name, email]);
    res.status(201).json({ status: 'created' });
});

app.delete('/users/:id', async (req: Request, res: Response) => {
    await pool.query('DELETE FROM users WHERE id = $1', [req.params.id]);
    res.status(204).send();
});

app.get('/search', (req: Request, res: Response) => {
    const q = req.query.q;
    res.send(`<h1>Results for ${q}</h1>`);
});

function startServer(port: number): void {
    app.listen(port, () => {
        console.log(`Server running on port ${port}`);
    });
}
