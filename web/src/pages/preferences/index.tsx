import type { NextPage } from 'next';
import Head from 'next/head';
import NextLink from 'next/link';
import {
  Box,
  Button,
  Container,
  FormControlLabel,
  Grid,
  Paper,
  Radio,
  RadioGroup,
  Typography
} from '@mui/material';
import { AuthGuard } from '../../components/authentication/auth-guard';
import { DashboardLayout } from '../../components/dashboard/dashboard-layout';

const typeOptions = [
  {
    description: 'Display your datas from your online AWS cloud provider',
    title: 'AWS',
    value: 'AWS'
  },
  {
    description: 'Display your datas from your online Scaleway cloud provider',
    title: 'Scaleway',
    value: 'Scaleway'
  }
];

const Preferences: NextPage = () => {
  return (
    <>
      <Head>
        <title>
          Preferences
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
                  Choose your preferences, AWS or Scaleway ?
                </Typography>
              </Grid>
            </Grid>
          </Box>
          <Grid
              item
              md={8}
              sm={12}
              xs={12}
            >
            <Grid
              container
              margin={3}
            >
              <RadioGroup
                  sx={{
                    '& > *:not(:last-child)': {
                      mb: 2
                    }
                  }}
                >
                  {typeOptions.map((typeOption, index) => (
                    <Paper
                      key={typeOption.value}
                      sx={{
                        alignItems: 'flex-start',
                        display: 'flex',
                        padding: 2
                      }}
                      variant="outlined"
                    >
                      <FormControlLabel
                        control={<Radio />}
                        disabled={index === 2 && true}
                        key={typeOption.value}
                        label={(
                          <Box sx={{ ml: 2 }}>
                            <Typography
                              sx={{
                                color: index === 2
                                  ? 'action.disabled'
                                  : 'text.primary'
                              }}
                              variant="subtitle2"
                            >
                              {typeOption.title}
                            </Typography>
                            <Typography
                              sx={{
                                color: index === 2
                                  ? 'action.disabled'
                                  : 'text.secondary'
                              }}
                              variant="body2"
                            >
                              {typeOption.description}
                            </Typography>
                          </Box>
                        )}
                        value={typeOption.value}
                      />
                    </Paper>
                  ))}
                </RadioGroup>
            </Grid>
            <NextLink
              href="/dashboard"
              passHref
            >
              <Button
                component="button"
                sx={{ mb: 3 }}
              >
                Validate your choice
              </Button>
            </NextLink>
          </Grid>
        </Container>
      </Box>
    </>
  );
};

Preferences.getLayout = (page) => (
  <AuthGuard>
    <DashboardLayout>
      {page}
    </DashboardLayout>
  </AuthGuard>
);

export default Preferences;
