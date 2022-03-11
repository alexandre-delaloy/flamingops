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
import { DotsHorizontal as DotsHorizontalIcon } from '../../../icons/dots-horizontal';
import { Mail as MailIcon } from '../../../icons/mail';
import { UserCircle as UserCircleIcon } from '../../../icons/user-circle';
import { Scrollbar } from '../../scrollbar';
import { SeverityPill } from '../../severity-pill';

const typeOptions = [
  {
    description: 'First available option',
    title: 'option 1',
    value: 'Option 1'
  },
  {
    description: 'Second available option',
    title: 'option 2',
    value: 'Option 2'
  },
  {
    description: 'Third available option',
    title: 'option 3',
    value: 'Option 3'
  },
  {
    description: 'Fourth available option',
    title: 'option 4',
    value: 'Option 4'
  }
];

export const AccountTeamSettings: FC = () => (

  
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
);
