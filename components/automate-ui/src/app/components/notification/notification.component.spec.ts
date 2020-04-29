import { CUSTOM_ELEMENTS_SCHEMA } from '@angular/core';
import { async, ComponentFixture, TestBed, tick, fakeAsync, flush } from '@angular/core/testing';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MockComponent } from 'ng2-mock-component';
import { using } from 'app/testing/spec-helpers';

import { ChefNotificationComponent } from './notification.component';
import { Type } from 'app/entities/notifications/notification.model';

describe('ChefNotificationComponent', () => {
  let component: ChefNotificationComponent;
  let fixture: ComponentFixture<ChefNotificationComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      imports: [ BrowserAnimationsModule ],
      declarations: [
        ChefNotificationComponent,
        MockComponent({ selector: 'chef-icon' })
      ],
      schemas: [ CUSTOM_ELEMENTS_SCHEMA ]
    })
      .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ChefNotificationComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  const genNotification = (type: Type, timeout: number) => {
    component.type = type;
    component.timeout = timeout;
  };

  describe('generates a notification', () => {

    it('and defaults to type of "info" and timeout of 8', () => {
      expect(component.type).toBe('info');
      expect(component.timeout).toEqual(8);
    });

    using([
      [Type.info, 7],
      [Type.info, 0],
      [Type.info, -1],
      [Type.error, 43],
      [Type.error, 0]
    ], function (notificationType: Type, timeLimit: number) {
        it(`and reads type of ${notificationType} and timeout of ${timeLimit} from inputs`, () => {
            genNotification(notificationType, timeLimit);

            expect(component.type).toBe(notificationType);
            expect(component.timeout).toEqual(timeLimit);
          });
        });
  });

  describe('After a notification is visible', () => {

    using([
      ['is not infinite', Type.info, 10, false],
      ['is infinite', Type.info, 0, true],
      ['is not infinite', Type.error, 43, false],
      ['is infinite', Type.error, 0, true],
      ['is infinite', Type.error, -1, true]
    ], function (description: string, type: Type, timeLimit: number, expected: boolean) {
      it(`${description} when type is ${type} and timeout is ${timeLimit}`, () => {
        genNotification(type, timeLimit);
        component.ngAfterViewInit();

        expect(component.isInfinite).toBe(expected);
      });
    });

    using([
      [Type.info, 1, false],
      [Type.info, 2, false],
      [Type.error, 3, false]
    ], function (type: Type, timeLimit: number, expected: boolean) {
      it(`notification closes when timeout is ${timeLimit}`, fakeAsync(() => {
        spyOn(component, 'handleClose');
        genNotification(type, timeLimit);
        component.ngAfterViewInit();

        expect(component.isInfinite).toBe(expected);
        tick((timeLimit * 1000) + 1);
        expect(component.handleClose).toHaveBeenCalled();

        fixture.destroy();
        flush();
      }));
    });

    using([
      [Type.info, 0, true],
      [Type.error, 0, true]
    ], function (type: Type, timeLimit: number, expected: boolean) {
      it(`${type} notification does not close when infinite`, fakeAsync(() => {
        spyOn(component, 'handleClose');
        genNotification(type, timeLimit);
        component.ngAfterViewInit();

        expect(component.isInfinite).toBe(expected);
        tick(500);
        expect(component.handleClose).not.toHaveBeenCalled();

        fixture.destroy();
        flush();
      }));
    });
  });

  describe('clicking the close button', () => {
    it('calls the handle close function', fakeAsync(() => {
      spyOn(component, 'handleClose');
      genNotification(Type.info, 2);
      component.ngAfterViewInit();

      const closeButton = fixture.debugElement.nativeElement.querySelector('chef-icon');
      closeButton.click();
      tick(500);
      expect(component.handleClose).toHaveBeenCalled();

      fixture.destroy();
      flush();
    }));
  });
});
