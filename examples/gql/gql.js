const { gql } = require("@apollo/client");

const TEST_QUERY = gql`
  query test {
    test
  }
`;

const REVERSE_MUTATION = gql`
  mutation reverse($str: String!) {
    reverse(str: $str) {
      string
      madeAt
    }
  }
`;

const REVERSED_SUB = gql`
    subscription stringReversed {
        stringReversed {
            madeAt
            string
        }
    }
`;

module.exports = {
  TEST_QUERY,
  REVERSE_MUTATION,
  REVERSED_SUB,
};
