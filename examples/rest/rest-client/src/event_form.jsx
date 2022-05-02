import React, { useEffect, useState } from "react";
import axios from "axios";
import { Button, Form, Input } from "antd";

export const EventForm = (props) => {
  const [title, setTitle] = useState("");
  const [description, setDescription] = useState("");
  const [appState, setAppState] = useState({
    loading: false,
    repos: null,
  });

  const createEvent = () => {
    setAppState({ loading: true, events: [] });
    const apiUrl = "http://localhost:3000/events.json";
    axios
      .post(apiUrl, {
        event: {
          title,
          description,
        },
      })
      .then((events) => {
        console.log(events);
        const allEvents = events.data;
        setAppState({ loading: false, events: allEvents });

        props.onEventCreated();
      })
      .catch((err) => {
        console.log(err);
        setAppState({ loading: false, events: [] });
      });
  };

  return (
    <Form
      style={{ width: "600px" }}
      labelCol={{ span: 8 }}
      wrapperCol={{ span: 16 }}
    >
      Event creation form:
      <br />
      <Form.Item label="Event title" wrapperCol={{ offset: 8, span: 16 }}>
        <Input value={title} onChange={(e) => setTitle(e.target.value)} />
      </Form.Item>
      <Form.Item label="Event description" wrapperCol={{ offset: 8, span: 16 }}>
        <Input
          value={description}
          onChange={(e) => setDescription(e.target.value)}
        />
      </Form.Item>
      <Form.Item wrapperCol={{ offset: 8, span: 16 }}>
        <Button onClick={createEvent}>Create</Button>
      </Form.Item>
    </Form>
  );
};
