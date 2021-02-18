import React, { useEffect, useState } from "react";
import axios from "axios";

export const EventList = (props) => {
  const [appState, setAppState] = useState({
    loading: false,
    repos: null,
  });

  useEffect(() => {
    setAppState({ loading: true, events: [] });
    const apiUrl = "http://localhost:3000/events.json";
    axios
      .get(apiUrl)
      .then((events) => {
        console.log(events);
        const allEvents = events.data;
        setAppState({ loading: false, events: allEvents });
      })
      .catch((err) => {
        console.log(err);
        setAppState({ loading: false, events: [] });
      });
  }, [setAppState, props.version]);

  return !!appState.loading ? (
    <div>Loading...</div>
  ) : (
    <div>
      Events:
      {(appState.events || []).map((event) => (
        <div>
          <b>{event.title}</b>
          {event.description}
        </div>
      ))}
    </div>
  );
};
