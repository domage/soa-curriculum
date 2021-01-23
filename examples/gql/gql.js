const { gql } = require('@apollo/client');

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

module.exports = {
    TEST_QUERY,
    REVERSE_MUTATION,
    REVERSED_SUB
}
