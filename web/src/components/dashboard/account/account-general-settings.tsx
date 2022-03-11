import type { FC } from 'react';
import {
  Avatar,
  Box,
  Button,
  Card,
  CardContent,
  Divider,
  FormControlLabel,
  Grid,
  Paper,
  Radio,
  RadioGroup,
  TextField,
  Typography
} from '@mui/material';
import { UserCircle as UserCircleIcon } from '../../../icons/user-circle';

export const AccountGeneralSettings: FC = (props) => {
  // To get the user from the authContext, you can use
  // `const { user } = useAuth();`
  const user = {
    avatar: '/static/mock-images/avatars/avatar-anika_visser.png',
    name: 'Anika Visser'
  };

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

  return (
    <Box
      sx={{ mt: 4 }}
      {...props}
    >
      
      <Card sx={{ mt: 4 }}>
        <CardContent>
          <Grid
            container
            spacing={3}
          >
            <Grid
              item
              md={4}
              xs={12}
            >
              <Typography variant="h6">
                Choice of provider
              </Typography>
            </Grid>
            <Grid
              item
              md={8}
              sm={12}
              xs={12}
            >
            <Divider />
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
          </Grid>
        </CardContent>
      </Card>
      <Card sx={{ mt: 4 }}>
        <CardContent>
          <Grid
            container
            spacing={3}
          >
            <Grid
              item
              md={4}
              xs={12}
            >
              <Typography variant="h6">
                Delete Account
              </Typography>
            </Grid>
            <Grid
              item
              md={8}
              xs={12}
            >
              <Typography
                sx={{ mb: 3 }}
                variant="subtitle1"
              >
                Delete your account and all of your source data. This is irreversible.
              </Typography>
              <Button
                color="error"
                variant="outlined"
              >
                Delete account
              </Button>
            </Grid>
          </Grid>
        </CardContent>
      </Card>
    </Box>
  );
};
