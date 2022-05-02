import React, { useState } from "react";
import logo from "./logo.svg";
import "./App.css";
import { EventList } from "./event_list";
import { EventForm } from "./event_form";

function App() {
  const [version, setVersion] = useState(0);

  console.log(version);

  return (
    <div className="App" style={{ width: '600px', margin: '0 auto' }}>
      App:
      <EventList version={version} />
      <EventForm onEventCreated={() => setVersion(version + 1)} />
    </div>
  );
}

export default App;
