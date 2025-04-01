"use client"

import React, { useState } from "react";
import type { GetTaskCandidateResponse, TaskCandidate } from "@/app/models/api";

export default function TaskCandidatePage() {
    const [text, setText] = useState('');
    const [taskCandidates, setTaskCandidates] = useState([] as TaskCandidate[]);

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

        const responseModel = {
            ...data
        } satisfies GetTaskCandidateResponse

        setTaskCandidates(responseModel.taskCandidates ?? [])
        console.log(taskCandidates)
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
            <ul className='md:flex  hidden flex-initial text-left'>
                {taskCandidates.map((value) => (
                    <li>
                        {`タスク名: ${value.name} 一致率: ${value.matchRate}`}
                    </li>
                ))}
            </ul>
        </div>
    )
}