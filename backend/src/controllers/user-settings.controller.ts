import { 
  Controller, 
  Post, 
  Get, 
  Put, 
  Delete, 
  Body, 
  Param, 
  HttpCode, 
  HttpStatus,
  ValidationPipe,
  UsePipes
} from '@nestjs/common';
import { 
  ApiTags, 
  ApiOperation, 
  ApiResponse, 
  ApiParam, 
  ApiBody,
  ApiCreatedResponse,
  ApiOkResponse,
  ApiNotFoundResponse,
  ApiBadRequestResponse
} from '@nestjs/swagger';
import { UserSettingsMockService } from '../services/user-settings-mock.service';
import { 
  SaveUserSettingsDto, 
  GetUserSettingsDto, 
  UserSettingsResponseDto,
  UpdateUserSettingsDto 
} from '../dto/user-settings.dto';

@ApiTags('User Settings')
@Controller('user-settings')
@UsePipes(new ValidationPipe({ transform: true }))
export class UserSettingsController {
  constructor(private readonly userSettingsService: UserSettingsMockService) {}

  @Post()
  @HttpCode(HttpStatus.CREATED)
  @ApiOperation({
    summary: 'Save or update user settings',
    description: 'Creates a new user settings record or updates existing one. Users are identified by their wallet address. Settings are flexible and can contain any user preferences.'
  })
  @ApiBody({ type: SaveUserSettingsDto })
  @ApiCreatedResponse({
    description: 'User settings saved successfully',
    type: UserSettingsResponseDto
  })
  @ApiBadRequestResponse({
    description: 'Invalid request data - wallet address or settings format is incorrect'
  })
  async saveUserSettings(@Body() saveUserSettingsDto: SaveUserSettingsDto): Promise<UserSettingsResponseDto> {
    return this.userSettingsService.saveUserSettings(saveUserSettingsDto);
  }

  @Get(':walletAddress')
  @ApiOperation({
    summary: 'Get user settings by wallet address',
    description: 'Retrieves user settings for the specified wallet address. Returns 404 if no settings found.'
  })
  @ApiParam({
    name: 'walletAddress',
    description: 'User wallet address (42 characters)',
    example: '0x1234567890abcdef1234567890abcdef12345678'
  })
  @ApiOkResponse({
    description: 'User settings retrieved successfully',
    type: UserSettingsResponseDto
  })
  @ApiNotFoundResponse({
    description: 'User settings not found for the specified wallet address'
  })
  async getUserSettings(@Param('walletAddress') walletAddress: string): Promise<UserSettingsResponseDto> {
    return this.userSettingsService.getUserSettings({ walletAddress });
  }

  @Put(':walletAddress')
  @ApiOperation({
    summary: 'Update user settings',
    description: 'Updates existing user settings by merging new settings with existing ones. Only updates the settings object, preserves other fields.'
  })
  @ApiParam({
    name: 'walletAddress',
    description: 'User wallet address (42 characters)',
    example: '0x1234567890abcdef1234567890abcdef12345678'
  })
  @ApiBody({ type: UpdateUserSettingsDto })
  @ApiOkResponse({
    description: 'User settings updated successfully',
    type: UserSettingsResponseDto
  })
  @ApiNotFoundResponse({
    description: 'User settings not found for the specified wallet address'
  })
  @ApiBadRequestResponse({
    description: 'Invalid settings format'
  })
  async updateUserSettings(
    @Param('walletAddress') walletAddress: string,
    @Body() updateSettingsDto: UpdateUserSettingsDto
  ): Promise<UserSettingsResponseDto> {
    return this.userSettingsService.updateUserSettings(walletAddress, updateSettingsDto.settings);
  }

  @Delete(':walletAddress')
  @HttpCode(HttpStatus.NO_CONTENT)
  @ApiOperation({
    summary: 'Delete user settings',
    description: 'Permanently deletes user settings for the specified wallet address. This action cannot be undone.'
  })
  @ApiParam({
    name: 'walletAddress',
    description: 'User wallet address (42 characters)',
    example: '0x1234567890abcdef1234567890abcdef12345678'
  })
  @ApiResponse({
    status: 204,
    description: 'User settings deleted successfully'
  })
  @ApiNotFoundResponse({
    description: 'User settings not found for the specified wallet address'
  })
  async deleteUserSettings(@Param('walletAddress') walletAddress: string): Promise<void> {
    return this.userSettingsService.deleteUserSettings(walletAddress);
  }
} 