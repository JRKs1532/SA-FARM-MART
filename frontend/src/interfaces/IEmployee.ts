import { EducationsInterface } from "./IEducation";
import { GendersInterface } from "./IGender";
import { PositionsInterface } from "./IPosition";


export interface EmployeeInterface {
  ID?: number;
  FirstName?: string;
  LastName?: string;
  Email?: string;
  Password?: string;
  Telephone?: string;
  Slary?: string;
  StartJob?: Date | null;
  EducationID?: number;
  Education?:EducationsInterface;
  GenderID?: number;
  Gender?: GendersInterface;
  PositionID?: number;
  Position?: PositionsInterface;
}