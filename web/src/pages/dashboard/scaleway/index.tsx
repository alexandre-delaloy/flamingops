import { useEffect } from 'react';
import type { NextPage } from 'next';
import Head from 'next/head';
import {
  Box,
  Container,
  Grid,
  Typography
} from '@mui/material';
// import { customerApi } from '../../../api/customer-api';
import { AuthGuard } from '../../../components/authentication/auth-guard';
import { DashboardLayout } from '../../../components/dashboard/dashboard-layout';
import { OverviewLatestTransactions } from '../../../components/dashboard/overview/overview-latest-transactions';
import { gtm } from '../../../lib/gtm';
// import type { Customer } from '../../../types/customer';

const CustomerList: NextPage = () => {

  useEffect(() => {
    gtm.push({ event: 'page_view' });
  }, []);

  return (
    <>
      <Head>
        <title>
          Dashboard: Scaleway List
        </title>
      </Head>
      <Box
        component="main"
        sx={{
          flexGrow: 1,
          py: 8
        }}
      >
        <Container maxWidth="xl">
          <Box sx={{ mb: 4 }}>
              <Grid
                container
                justifyContent="space-between"
                spacing={3}
              >
                <Grid item>
                  <Typography variant="h4">
                    Scaleway
                  </Typography>
                </Grid>
              </Grid>
            </Box>
          <OverviewLatestTransactions />
        </Container>
      </Box>
    </>
  );
};

CustomerList.getLayout = (page) => (
  <AuthGuard>
    <DashboardLayout>
      {page}
    </DashboardLayout>
  </AuthGuard>
);

export default CustomerList;
