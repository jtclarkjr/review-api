import { PrismaClient } from '@prisma/client';
import * as bcrypt from 'bcrypt';

const prisma = new PrismaClient();

async function main() {
  const passwordHash = await bcrypt.hash('adminpassword', 10);

  // Create the admin user
  const admin = await prisma.employee.upsert({
    where: { email: 'admin@company.com' },
    update: {},
    create: {
      name: 'Admin User',
      email: 'admin@company.com',
      password: passwordHash,
      isAdmin: true,
    },
  });

  // Create other employees
  const employees = [
    {
      name: 'John Doe',
      email: 'john.doe@company.com',
      password: await bcrypt.hash('password123', 10),
      isAdmin: false,
    },
    {
      name: 'Jane Smith',
      email: 'jane.smith@company.com',
      password: await bcrypt.hash('password456', 10),
      isAdmin: false,
    },
    {
      name: 'Emily Johnson',
      email: 'emily.johnson@company.com',
      password: await bcrypt.hash('password789', 10),
      isAdmin: false,
    },
  ];

  for (const employee of employees) {
    await prisma.employee.upsert({
      where: { email: employee.email },
      update: {},
      create: employee,
    });
  }

  console.log({ admin, employees });
}

main()
  .catch((e) => {
    console.error(e);
    process.exit(1);
  })
  .finally(async () => {
    await prisma.$disconnect();
  });
