import { Box, Typography } from "@material-ui/core";

import Layout, { Page } from "../../components/Layout";

function Index(): JSX.Element {
  return (
    <Layout page={Page.Projects} title="Projects">
      <Box p={4}>
        <Typography paragraph>
          Projects contain settings and data generated/processed by Hetty. They
          are stored as SQLite database files on disk.
        </Typography>
      </Box>
    </Layout>
  );
}

export default Index;
