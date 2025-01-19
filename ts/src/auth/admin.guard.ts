import {
  Injectable,
  CanActivate,
  ExecutionContext,
  UnauthorizedException,
} from '@nestjs/common';
import { JwtService } from '@nestjs/jwt';
import { PrismaService } from 'prisma/prisma.service';
import { JwtPayload } from './jwt-payload.interface';
import { TokenExpiredError } from 'jsonwebtoken';

@Injectable()
export class AdminGuard implements CanActivate {
  constructor(
    private jwtService: JwtService,
    private prisma: PrismaService,
  ) {}

  async canActivate(context: ExecutionContext): Promise<boolean> {
    const request = context.switchToHttp().getRequest();
    const authHeader = request.headers.authorization;

    if (!authHeader) {
      throw new UnauthorizedException('Authorization token not found');
    }

    const token = authHeader.split(' ')[1];
    try {
      const decoded: JwtPayload = this.jwtService.verify(token);

      const user = await this.prisma.employee.findUnique({
        where: { id: decoded.sub },
      });

      if (!user || !user.isAdmin) {
        throw new UnauthorizedException('User not found or not an admin');
      }

      request.user = user; // Attach the user to the request
      return true;
    } catch (err) {
      if (err instanceof TokenExpiredError) {
        throw new UnauthorizedException('Token expired');
      }
      throw new UnauthorizedException('Invalid token or user not found');
    }
  }
}
