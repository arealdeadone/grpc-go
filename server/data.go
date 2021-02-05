package main

import (
  "grpc-go/pb"
)

var employees = []pb.Employee {
  pb.Employee {
    Id: 1,
    BadgeNumber: 1242813,
    FirstName: "Arvind",
    LastName: "Rachuri",
    VacationAccrualRate: 1.5,
    VacationAccrued: 14,
  },
  pb.Employee {
    Id: 2,
    BadgeNumber: 1242814,
    FirstName: "Parag",
    LastName: "Jain",
    VacationAccrualRate: 2,
    VacationAccrued: 30,
  },
  pb.Employee {
    Id: 3,
    BadgeNumber: 1242815,
    FirstName: "Pulkit",
    LastName: "Rathi",
    VacationAccrualRate: 1.5,
    VacationAccrued: 14,
  },
}
