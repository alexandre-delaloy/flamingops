import type { NextPage } from 'next';
import Head from 'next/head';
import {
  Box,
  Container,
  Grid,
  Typography
} from '@mui/material';
import { AuthGuard } from '../components/authentication/auth-guard';
import { DashboardLayout } from '../components/dashboard/dashboard-layout';
import { OverviewCryptoWallet } from '../components/dashboard/overview/overview-crypto-wallet';
import { OverviewLatestTransactions } from '../components/dashboard/overview/overview-latest-transactions';
import { OverviewPrivateWallet } from '../components/dashboard/overview/overview-private-wallet';
import { OverviewTotalBalance } from '../components/dashboard/overview/overview-total-balance';
import { OverviewTotalTransactions } from '../components/dashboard/overview/overview-total-transactions';

const Overview: NextPage = () => {
  return (
    <>
      <Head>
        <title>
          Dashboard: Overview
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
                  Good Morning
                </Typography>
              </Grid>
            </Grid>
          </Box>
          <Grid
            container
            spacing={4}
          >
            <Grid
              item
              md={6}
              xs={12}
            >
              <OverviewCryptoWallet />
            </Grid>
            <Grid
              item
              md={6}
              xs={12}
            >
              <OverviewPrivateWallet />
            </Grid>
            <Grid
              item
              md={8}
              xs={12}
            >
              <OverviewTotalTransactions />
            </Grid>
            <Grid
              item
              md={4}
              xs={12}
            >
              <OverviewTotalBalance />
            </Grid>
            <Grid
              item
              md={8}
              xs={12}
            >
              <OverviewLatestTransactions />
            </Grid>
          </Grid>
        </Container>
      </Box>
    </>
  );
};

Overview.getLayout = (page) => (
  <AuthGuard>
    <DashboardLayout>
      {page}
    </DashboardLayout>
  </AuthGuard>
);

export default Overview;
