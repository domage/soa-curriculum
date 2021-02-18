import React, { useEffect, useState } from "react";
import axios from "axios";

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
    <div>
      Event creation form:
      <br />
      <input value={title} onChange={(e) => setTitle(e.target.value)} />
      <input
        value={description}
        onChange={(e) => setDescription(e.target.value)}
      />
      <button onClick={createEvent}>Create</button>
    </div>
  );
};
