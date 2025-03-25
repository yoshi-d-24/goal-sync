"use client"

import React, { useState } from "react";

export default function TaskCandidate() {
    const [text, setText] = useState('');
    const [candidates, setCandidates] = useState('');

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault()
        const res = await fetch('http://localhost:8080/task-candidates',
            { 
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    text: text,
                    job: 'ITエンジニア'
                })
            });

        if (!res.ok) {
            console.log('failed to get candidates')
        }
        const data = await res.json();
        
        console.log(data);
        setCandidates(JSON.parse(data));
    }

    return (
        <div>
            <h1>タスク候補リストアップ</h1>
            <form onSubmit={handleSubmit}>
                <textarea
                    value={text}
                    onChange={(e) => { setText(e.target.value) }}
                ></textarea>
                <button type="submit">Button</button>
            </form>
            <span content={candidates}></span>
        </div>
    )
}