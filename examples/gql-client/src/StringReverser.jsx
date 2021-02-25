import React, { useEffect, useState } from "react";
import { useQuery, useMutation } from "@apollo/client";
import { Button, Input, Spin, message } from "antd";

import { TEST_QUERY, REVERSE, REVERSED_SUBSCRIPTION } from "./gql";

const Sub = ({ sub }) => {
  useEffect(() => sub(), []);

  return <div />;
};

export const StringReserver = () => {
  const [string, setString] = useState("");

  const { loading, error, data, subscribeToMore } = useQuery(TEST_QUERY);
  const [reverseMutation] = useMutation(REVERSE);

  return (
    <div
      style={{
        width: "600px",
        margin: "0 auto",
        padding: "1rem",
        textAlign: "center",
        verticalAlign: "center",
      }}
    >
      {!!loading ? (
        <Spin size="large" style={{ margin: "0 auto" }} />
      ) : (
        `Query result: ${data.test}`
      )}
      <Sub
        sub={() =>
          subscribeToMore({
            document: REVERSED_SUBSCRIPTION,
            variables: null,
            updateQuery: (prev, data) => {
                console.log(data);
                message.info(data.subscriptionData.data.stringReversed.string);
            },
          })
        }
      />
      <Input value={string} onChange={(e) => setString(e.target.value)} />
      <Button onClick={() => reverseMutation({ variables: { str: string } })}>
        Reverse string
      </Button>
    </div>
  );
};
