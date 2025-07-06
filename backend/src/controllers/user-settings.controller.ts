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
import { UserSettingsService } from '../services/user-settings.service';
import { SaveUserSettingsDto, GetUserSettingsDto, UserSettingsResponseDto } from '../dto/user-settings.dto';

@Controller('user-settings')
@UsePipes(new ValidationPipe({ transform: true }))
export class UserSettingsController {
  constructor(private readonly userSettingsService: UserSettingsService) {}

  @Post()
  @HttpCode(HttpStatus.CREATED)
  async saveUserSettings(@Body() saveUserSettingsDto: SaveUserSettingsDto): Promise<UserSettingsResponseDto> {
    return this.userSettingsService.saveUserSettings(saveUserSettingsDto);
  }

  @Get(':walletAddress')
  async getUserSettings(@Param('walletAddress') walletAddress: string): Promise<UserSettingsResponseDto> {
    return this.userSettingsService.getUserSettings({ walletAddress });
  }

  @Put(':walletAddress')
  async updateUserSettings(
    @Param('walletAddress') walletAddress: string,
    @Body() settings: Record<string, any>
  ): Promise<UserSettingsResponseDto> {
    return this.userSettingsService.updateUserSettings(walletAddress, settings);
  }

  @Delete(':walletAddress')
  @HttpCode(HttpStatus.NO_CONTENT)
  async deleteUserSettings(@Param('walletAddress') walletAddress: string): Promise<void> {
    return this.userSettingsService.deleteUserSettings(walletAddress);
  }
} 