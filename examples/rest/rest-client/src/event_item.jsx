import React, { useEffect, useState } from "react";
import axios from "axios";
import { Button, List, Modal, Input, Form } from "antd";

export const EventItem = (props) => {
  const [appState, setAppState] = useState({
    loading: false,
    repos: null,
  });
  const [visible, setVisible] = useState(false);
  const [title, setTitle] = useState("");
  const [description, setDescription] = useState("");
  const [vers, setVers] = useState(0);

  useEffect(() => {
    setAppState({ loading: true, comments: [] });
    const apiUrl =
      "http://localhost:3000/events/" + props.event.id + "/comments.json";
    axios
      .get(apiUrl)
      .then((comments) => {
        const allComments = comments.data;
        setAppState({ loading: false, comments: allComments });
      })
      .catch((err) => {
        setAppState({ loading: false, comments: [] });
      });
  }, [setAppState, props.version, vers]);

  const createComment = () => {
    setAppState({ loading: true, events: [] });
    const apiUrl =
      "http://localhost:3000/events/" + props.event.id + "/comments.json";
    axios
      .post(apiUrl, {
        comment: {
          title,
          description,
        },
      })
      .then((events) => {
        setVers(vers + 1);
      });
  };

  if (appState.loading) {
    return "Loading...";
  }

  return (
    <List.Item
      actions={[<Button onClick={() => setVisible(true)}>Add comment</Button>]}
    >
      <Modal
        visible={visible}
        onOk={() => {
          createComment();
          setVisible(false);
        }}
        onCancel={() => setVisible(false)}
      >
        <Form
          labelCol={{ span: 4 }}
          wrapperCol={{ span: 20 }}
        >
          Comment creation form:
          <br />
          <Form.Item label="Comment title" wrapperCol={{ offset: 8, span: 16 }}>
            <Input value={title} onChange={(e) => setTitle(e.target.value)} />
          </Form.Item>
          <Form.Item
            label="Comment description"
            wrapperCol={{ offset: 8, span: 16 }}
          >
            <Input
              value={description}
              onChange={(e) => setDescription(e.target.value)}
            />
          </Form.Item>
        </Form>
      </Modal>
      <List.Item.Meta
        title={props.event.title}
        description={
          <div>
            <div>{props.event.description}</div>
            Comments:
            {(appState.comments || []).map((comment) => (
              <div>
                <b>{comment.title}</b> {comment.description}
              </div>
            ))}
          </div>
        }
      />
    </List.Item>
  );
};
